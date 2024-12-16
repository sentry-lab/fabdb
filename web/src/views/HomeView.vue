<script setup lang="ts">
import Column from "primevue/column";
import DataTable from "primevue/datatable";
import { onMounted, ref } from "vue";

import { GetWorkflows, type Workflow } from "@/services/workflows";

const workflows = ref<Workflow[]>([]);

onMounted(async () => {
  workflows.value = await GetWorkflows();
});
</script>

<template>
  <h1>Hello!</h1>
  <DataTable :value="workflows">
    <Column field="id" header="ID" />
    <Column field="name" header="Name" />
  </DataTable>
</template>
