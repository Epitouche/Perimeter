export interface Area {
  id: number;
  createdAt: string;
  action: {
    name: string;
    service: {
      name: string;
    };
  };
  reaction: {
    name: string;
    service: {
      name: string;
    };
  };
}
