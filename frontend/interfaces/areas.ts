export interface Area {
  id: number;
  createdAt: string;
  title: string;
  description: string;
  action_refresh_rate: number;
  enable: boolean;
  action: {
    id: number;
    name: string;
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
    service: {
      name: string;
      color: string;
      icon: string;
    };
  };
  reaction_option?: string | object;
}
