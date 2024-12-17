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

export interface BlockTypeProperty {
  id: string;
  name: string;
  value_type: string;
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
            id: "4",
            name: "wow",
            value_type: "text",
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
          id: "3",
          name: "wow",
          value: "hello",
        },
      ],
    },
    {
      id: "2",
      position: 1,
      workflow_id: workflow_id,
      block_type: {
        id: "2",
        name: "Other Block",
        is_preset: false,
        properties: [
          {
            id: "2",
            name: "Worked on",
            value_type: "text",
            default_value: "",
          },
        ],
      },
      properties: [
        {
          id: "2",
          name: "Worked on",
          value: "hello",
        },
      ],
    },
  ];
};
