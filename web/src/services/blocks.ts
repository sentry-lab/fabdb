export interface Block {
  id: string;
  position: number;
  workflow_id: string;
  block_type: BlockType;
  properties: BlockProperty[];
}

export interface BlockProperty {
  id: string;
  name: string;
  value: string;
}

export interface BlockType {
  id: string;
  name: string;
  is_preset: boolean;
  properties: BlockTypeProperty[];
}

export type ValueType = "text" | "boolean" | "number" | "date";
export interface BlockTypeProperty {
  id: string;
  name: string;
  value_type: ValueType;
  default_value: string;
}

export const GetBlocks = async (workflow_id: string): Promise<Block[]> => {
  return [
    {
      id: "1",
      position: 0,
      workflow_id: workflow_id,
      block_type: {
        id: "1",
        name: "Test Block",
        is_preset: false,
        properties: [
          {
            id: "1",
            name: "test",
            value_type: "number",
            default_value: "",
          },
          {
            id: "2",
            name: "wow",
            value_type: "text",
            default_value: "",
          },
          {
            id: "3",
            name: "Happened on",
            value_type: "date",
            default_value: "",
          },
          {
            id: "4",
            name: "Complete",
            value_type: "boolean",
            default_value: "",
          },
        ],
      },
      properties: [
        {
          id: "1",
          name: "test",
          value: "123",
        },
        {
          id: "2",
          name: "wow",
          value: "hello",
        },
        {
          id: "3",
          name: "Happened on",
          value: "2024-12-01",
        },
        {
          id: "4",
          name: "Complete",
          value: "true",
        },
      ],
    },
  ];
};

export const SaveBlocks = async (blocks: Block[]) => {
  console.log("saving", JSON.stringify(blocks));
};
