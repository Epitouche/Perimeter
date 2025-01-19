/**
 * Represents a type object (action or reaction)
 * @property {number} id - The id of the type.
 * @property {string} name - The name of the type.
 * @property {string} description - The description of the type.
 * @property {string | object} option - The option for the type.
 */
export interface Type {
  id: number;
  name: string;
  description: string;
  option?: string | object;
}
