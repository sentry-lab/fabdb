<script setup lang="ts">
import Menubar from "primevue/menubar";
import type { MenuItem } from "primevue/menuitem";
import { ref } from "vue";

import Account from "@/icons/Account.vue";
import router from "@/router";

const items = ref<MenuItem[]>([
  {
    label: "Workflows",
    command: () => {
      router.push({ name: "workflows" });
    },
  },
  {
    label: "Wafers",
    command: () => {},
  },
  {
    label: "Chips",
    command: () => {},
  },
]);
</script>

<template>
  <div style="margin: 0 1em">
    <Menubar :model="items" style="margin: 0.75em 0">
      <template #end><Account style="width: 24px; height: 24px" /></template>
    </Menubar>
    <div style="margin: 0.75em 0">
      <slot></slot>
    </div>
  </div>
</template>
