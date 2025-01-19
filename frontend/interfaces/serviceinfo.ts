/**
 * Represents a service with its information, with the following properties:
 * @property {number} id - The ID of the service.
 * @property {string} name - The name of the service.
 * @property {string} color - The color of the service.
 * @property {string} icon - The icon of the service.
 * @property {string} description - The description of the service.
 * @property {boolean} oauth - Whether the service uses OAuth.
 */
export interface ServiceInfo {
  id: number;
  name: string;
  color: string;
  icon: string;
  description: string;
  oauth: boolean;
}
