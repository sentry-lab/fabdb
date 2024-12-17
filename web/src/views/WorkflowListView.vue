<script setup lang="ts">
import Column from "primevue/column";
import DataTable from "primevue/datatable";
import { onMounted, ref } from "vue";
import { RouterLink } from "vue-router";

import { GetWorkflows, type Workflow } from "@/services/workflows";

const workflows = ref<Workflow[]>([]);

onMounted(async () => {
  workflows.value = await GetWorkflows();
});
</script>

<template>
  <DataTable :value="workflows">
    <Column header="Name">
      <template #body="slotProps">
        <RouterLink
          :to="{
            name: 'workflow',
            params: {
              id: slotProps.data.id,
            },
          }"
        >
          {{ slotProps.data.name }}
        </RouterLink>
      </template>
    </Column>
  </DataTable>
</template>
