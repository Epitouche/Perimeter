/**
 * Represents an Area object with the following properties
 * @property {number} id - The unique identifier of the area
 * @property {string} createdAt - The date and time the area was created
 * @property {string} title - The title of the area
 * @property {string} description - The description of the area
 * @property {number} action_refresh_rate - The refresh rate of the area
 * @property {boolean} enable - The status of the area
 * 
 * @property {object} action - The action object of the area
 * @property {number} action.id - The unique identifier of the action
 * @property {string} action.name - The name of the action
 * @property {object} action.service - The service object of the action
 * @property {string} action.service.name - The name of the service
 * @property {string} action.service.color - The color of the service
 * @property {string} action.service.icon - The icon of the service
 * 
 * @property {string | object} action_option - The option of the action
 * 
 * @property {object} reaction - The reaction object of the area
 * @property {number} reaction.id - The unique identifier of the reaction
 * @property {string} reaction.name - The name of the reaction
 * @property {object} reaction.service - The service object of the reaction
 * @property {string} reaction.service.name - The name of the service
 * @property {string} reaction.service.color - The color of the service
 * @property {string} reaction.service.icon - The icon of the service
 * 
 * @property {string | object} reaction_option - The option of the reaction
 */
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
