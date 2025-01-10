export interface Service {
    name: string;
    oauth: boolean;
  }
  
  export interface Token {
    service: Service;
    id: number;
  }
  
  export interface ServiceResponse {
    tokens: Token[];
  }
  
