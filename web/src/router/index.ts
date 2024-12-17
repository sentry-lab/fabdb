import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: () => import("@/views/HomeView.vue"),
    },
    {
      path: "/workflows",
      name: "workflows",
      component: () => import("@/views/WorkflowListView.vue"),
    },
    {
      path: "/workflows/:id",
      name: "workflow",
      component: () => import("@/views/WorkflowView.vue"),
    },
  ],
});

export default router;
