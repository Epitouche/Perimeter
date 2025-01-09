export interface Service {
    name: string;
    id: number;
  }
  
  export interface Token {
    service: Service;
  }
  
  export interface ServiceResponse {
    tokens: Token[];
  }
  