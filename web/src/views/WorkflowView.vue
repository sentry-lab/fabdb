<script setup lang="ts">
import Button from "primevue/button";
import Column from "primevue/column";
import DataTable, { type DataTableRowReorderEvent } from "primevue/datatable";
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";

import ContentSave from "@/icons/ContentSave.vue";
import Delete from "@/icons/Delete.vue";
import Plus from "@/icons/Plus.vue";
import PuzzleEdit from "@/icons/PuzzleEdit.vue";
import PuzzleEditOutline from "@/icons/PuzzleEditOutline.vue";
import { type Block, GetBlocks } from "@/services/blocks";
import { GetWorkflow, type Workflow } from "@/services/workflows";

const route = useRoute();
const workflow = ref<Workflow>();
const blocks = ref<Block[]>();
const currentlyEditing = ref<number>(0);

const save = () => {
  console.log("save");
};

const addBlock = () => {
  console.log("add block");
};

const deleteBlock = () => {
  console.log("delete block", currentlyEditing.value);
};

const editBlock = (index: number) => {
  currentlyEditing.value = index;
};

const onRowReorder = (e: DataTableRowReorderEvent) => {
  let currentId = blocks.value![currentlyEditing.value].id;
  blocks.value = e.value;
  currentlyEditing.value = blocks.value.findIndex((b) => b.id == currentId);
};

onMounted(async () => {
  workflow.value = await GetWorkflow(route.params.id as string);
  blocks.value = await GetBlocks(workflow.value.id);
});
</script>

<template>
  <div style="margin-bottom: 2em">
    <h1>{{ workflow?.name }}</h1>
    Workflow
  </div>

  <div>
    <Button @click="addBlock()"><Plus style="height: 1.5em" />Add Block</Button>
    <Button
      @click="save()"
      style="margin-left: 1em; background-color: var(--p-gray-500); border-color: var(--p-gray-500)"
    >
      <ContentSave style="height: 1.5em" />Save
    </Button>
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
    <div v-if="blocks" style="flex: 8; display: flex; flex-direction: column; margin: 0 2em">
      <div>
        <h2 style="display: inline">{{ blocks[currentlyEditing].block_type.name }}</h2>
        <div style="float: right">
          <Button
            @click="deleteBlock()"
            style="background-color: var(--p-red-500); border-color: var(--p-red-500)"
            ><Delete style="height: 1.5em" />Delete Block</Button
          >
        </div>
      </div>
      <div v-for="prop in blocks[currentlyEditing].properties" :key="prop.id" style="display: flex">
        <div style="flex: 3">
          <b>{{ prop.name }}</b> ({{
            blocks[currentlyEditing].block_type.properties.find((p) => p.name == prop.name)
              ?.value_type
          }})
        </div>
        <div style="flex: 6">
          {{ prop.value }}
        </div>
      </div>
    </div>
  </div>
</template>
