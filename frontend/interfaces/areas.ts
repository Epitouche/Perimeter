export interface Area {
  id: number;
  createdAt: string;
  action: {
    name: string;
    service: {
      name: string;
      color: string;
      icon: string;
    };
  };
  reaction: {
    name: string;
    service: {
      name: string;
      color: string;
      icon: string;
    };
  };
}
