export interface Workflow {
  id: string;
  name: string;
}

export const GetWorkflows = async (): Promise<Workflow[]> => {
  return [
    { id: "1", name: "Workflow 1" },
    { id: "2", name: "Workflow 2" },
  ];
};

export const GetWorkflow = async (id: string): Promise<Workflow> => {
  if (id == "1") {
    return { id: "1", name: "Workflow 1" };
  }
  return { id: "2", name: "Workflow 2" };
};
