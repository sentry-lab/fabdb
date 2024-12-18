<script setup lang="ts">
import Button from "primevue/button";
import Column from "primevue/column";
import DataTable, { type DataTableRowReorderEvent } from "primevue/datatable";
import Popover from "primevue/popover";
import Select from "primevue/select";
import { onMounted, reactive, ref, watch } from "vue";
import { useRoute } from "vue-router";

import SmartInput from "@/components/SmartInput.vue";
import Delete from "@/icons/Delete.vue";
import Plus from "@/icons/Plus.vue";
import PuzzleEdit from "@/icons/PuzzleEdit.vue";
import PuzzleEditOutline from "@/icons/PuzzleEditOutline.vue";
import {
  type Block,
  type BlockType,
  GetBlockTypes,
  GetBlocks,
  SaveBlocks,
} from "@/services/blocks";
import { GetWorkflow, type Workflow } from "@/services/workflows";

const route = useRoute();
const addPopover = ref<InstanceType<typeof Popover> | null>();

const workflow = ref<Workflow>();
const blocks = reactive<Block[]>([]);
const blockTypes = ref<BlockType[]>();

const currentlyEditing = ref<number>(0);
const saving = ref(false);
const saveRemaing = ref(false);

const blockToAdd = ref<BlockType>();

watch(blocks, async (newBlock, _) => {
  saveRemaing.value = true;
  if (saving.value) {
    return;
  }
  saving.value = true;
  while (saveRemaing.value) {
    saveRemaing.value = false;
    await SaveBlocks(newBlock);
    await new Promise((r) => setTimeout(r, 1000));
  }
  saving.value = false;
});

const addBlock = () => {
  addPopover.value!.hide();
  let newBlock: Block = {
    id: "",
    position: blocks.length,
    workflow_id: workflow.value!.id,
    block_type: blockToAdd.value!,
    properties: [],
  };
  for (let i = 0; i < blockToAdd.value!.properties.length; i++) {
    newBlock.properties.push({
      id: "",
      name: blockToAdd.value!.properties[i].name,
      value: blockToAdd.value!.properties[i].default_value,
    });
  }
  blocks.push(newBlock);

  blockToAdd.value = undefined;
};

const deleteBlock = () => {
  blocks.splice(currentlyEditing.value, 1);
};

const editBlock = (index: number) => {
  currentlyEditing.value = index;
};

const onRowReorder = (e: DataTableRowReorderEvent) => {
  let currentId = blocks[currentlyEditing.value].id;
  for (let i = 0; i < blocks.length; i++) {
    blocks[i] = e.value[i];
  }
  currentlyEditing.value = blocks.findIndex((b) => b.id == currentId);
};

onMounted(async () => {
  workflow.value = await GetWorkflow(route.params.id as string);
  let unReactiveBlocks = await GetBlocks(workflow.value.id);
  for (let i = 0; i < unReactiveBlocks.length; i++) {
    blocks.push(unReactiveBlocks[i]);
  }
  blockTypes.value = await GetBlockTypes();
});
</script>

<template>
  <div style="margin-bottom: 2em">
    <h1>{{ workflow?.name }}</h1>
    Workflow
  </div>

  <div>
    <Button
      @click="
        (e) => {
          addPopover!.toggle(e);
        }
      "
    >
      <Plus style="height: 1.5em" />Add Block
    </Button>
    <Popover ref="addPopover">
      <Select
        v-model="blockToAdd"
        :options="blockTypes"
        optionLabel="name"
        style="width: 15em"
        @change="addBlock"
      />
    </Popover>

    <span style="margin-left: 1em; color: var(--p-gray-500)">
      <span v-if="saving">Saving...</span><span v-else>Saved!</span>
    </span>
  </div>
  <div style="display: flex">
    <div style="flex: 4">
      <DataTable :value="blocks" @rowReorder="onRowReorder">
        <Column rowReorder />
        <Column field="block_type.name" />
        <Column>
          <template #body="slotProps">
            <div @click="editBlock(slotProps.index)" style="cursor: pointer">
              <PuzzleEdit v-if="currentlyEditing == slotProps.index" style="height: 1.5em" />
              <PuzzleEditOutline v-else style="height: 1.5em" />
            </div>
          </template>
        </Column>
      </DataTable>
    </div>
    <div
      v-if="blocks.length > 0"
      style="flex: 8; display: flex; flex-direction: column; margin: 0 2em"
    >
      <div>
        <h2 style="display: inline">{{ blocks[currentlyEditing].block_type.name }}</h2>
        <div style="float: right; margin-bottom: 0.5em">
          <Button @click="deleteBlock()" severity="danger">
            <Delete style="height: 1.5em" />Delete Block
          </Button>
        </div>
      </div>
      <div
        v-for="i in blocks[currentlyEditing].properties.length"
        style="display: flex; margin: 0.25em 0; align-items: center"
      >
        <div style="flex: 6">
          <b>{{ blocks[currentlyEditing].properties[i - 1].name }}</b>
        </div>
        <div style="flex: 6">
          <SmartInput
            v-model="blocks[currentlyEditing].properties[i - 1].value"
            :value-type="blocks[currentlyEditing].block_type.properties[i - 1].value_type"
            :default-value="blocks[currentlyEditing].block_type.properties[i - 1].default_value"
          />
        </div>
      </div>
    </div>
  </div>
</template>
