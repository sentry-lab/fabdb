import { definePreset } from "@primevue/themes";
import Aura from "@primevue/themes/aura";
import { createPinia } from "pinia";
import PrimeVue from "primevue/config";
import { createApp } from "vue";

import App from "@/App.vue";
import "@/assets/styles/main.css";
import router from "@/router";

const app = createApp(App);

app.use(createPinia());
app.use(router);

const PrimeVuePreset = definePreset(Aura, {
  semantic: {
    primary: {
      50: "{purple.50}",
      100: "{purple.100}",
      200: "{purple.200}",
      300: "{purple.300}",
      400: "{purple.400}",
      500: "{purple.500}",
      600: "{purple.600}",
      700: "{purple.700}",
      800: "{purple.800}",
      900: "{purple.900}",
      950: "{purple.950}",
    },
  },
});

app.use(PrimeVue, {
  theme: {
    preset: PrimeVuePreset,
    options: {
      darkModeSelector: ".dark-mode",
    },
  },
  ripple: true,
});

app.mount("#app");
