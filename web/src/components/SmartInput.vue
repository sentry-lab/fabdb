<script setup lang="ts">
import Checkbox from "primevue/checkbox";
import DatePicker from "primevue/datepicker";
import InputNumber from "primevue/inputnumber";
import InputText from "primevue/inputtext";
import { computed } from "vue";

import type { ValueType } from "@/services/blocks";

const model = defineModel<string>();
const props = defineProps<{
  valueType: ValueType;
  defaultValue: string;
}>();

const usableModel = computed<any>({
  get() {
    let val = model.value || props.defaultValue;
    switch (props.valueType) {
      case "text":
        return val;
      case "number":
        if (val == "") {
          return null;
        }
        return parseInt(val);
      case "boolean":
        return val == "true";
      case "date":
        if (val != "") {
          return new Date(val + "T00:00");
        }
        return null;
    }
  },
  set(newValue) {
    switch (props.valueType) {
      case "text":
        model.value = newValue;
        break;
      case "number":
        if (newValue == null) {
          model.value = "";
          break;
        }
      case "boolean":
        model.value = newValue.toString();
        break;
      case "date":
        model.value = `${newValue.getFullYear()}-${newValue.getMonth() + 1}-${("0" + newValue.getDate()).slice(-2)}`;
        break;
    }
  },
});
</script>

<template>
  <InputText v-if="valueType == 'text'" type="text" v-model="usableModel" fluid />
  <InputNumber v-else-if="valueType == 'number'" v-model="usableModel" fluid />
  <Checkbox v-else-if="valueType == 'boolean'" v-model="usableModel" binary />
  <DatePicker
    v-else-if="valueType == 'date'"
    v-model="usableModel"
    date-format="yy-mm-dd"
    showIcon
    iconDisplay="input"
    :showButtonBar="true"
    fluid
  />
</template>
