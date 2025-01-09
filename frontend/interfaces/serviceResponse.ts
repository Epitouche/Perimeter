export interface Service {
  name: string;
}

export interface Token {
  service: Service;
}

export interface ServiceResponse {
  tokens: Token[];
}
