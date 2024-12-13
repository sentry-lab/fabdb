import api from "@/api";

export interface Workflow {
  id: string;
  name: string;
}

export const GetWorkflows = async (): Promise<Workflow[]> => {
  let resp = await api.get("/workflows");

  return resp.data;
};
