export interface Area {
  id: number;
  createdAt: string;
  title: string;
  description: string;
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
