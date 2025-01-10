export interface Service {
    name: string;
    id: number;
    oauth: boolean;
  }
  
  export interface Token {
    service: Service;
  }
  
  export interface ServiceResponse {
    tokens: Token[];
  }
  
