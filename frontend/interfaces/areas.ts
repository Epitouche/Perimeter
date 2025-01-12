export interface Area {
  id: number;
  createdAt: string;
  title: string;
  description: string;
  action: {
    id: number;
    name: string;
    option?: string | object;
    service: {
      name: string;
      color: string;
      icon: string;
    };
  };
  action_option?: string | object;
  reaction: {
    id: number;
    name: string;
    option?: string | object;
    service: {
      name: string;
      color: string;
      icon: string;
    };
  };
  reaction_option?: string | object;
}
