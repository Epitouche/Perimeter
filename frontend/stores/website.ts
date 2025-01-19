/**
 * This is used to store the state of the website, specifically for the workflow page.
 *
 * The store contains the following state:
 * @property {boolean} showNavBar - Whether the navigation bar should be shown.
 * @property {boolean} showCancelButton - Whether the cancel button should be shown.
 * @property {boolean} reactionButtonisDisabled - Whether the reaction button should be disabled.
 * @property {boolean} showCreateButton - Whether the create button should be shown.
 * @property {boolean} actionIsSelected - Whether an action has been selected.
 * @property {boolean} reactionIsSelected - Whether a reaction has been selected.
 * @property {string | null} actionId - The ID of the selected action.
 * @property {Record<string, unknown>} actionOptions - The options for the selected action.
 * @property {string | null} reactionId - The ID of the selected reaction.
 * @property {Record<string, unknown>} reactionOptions - The options for the selected reaction.
 * @property {string | null} actionServiceId - The ID of the service of the selected action.
 * @property {string | null} reactionServiceId - The ID of the service of the selected reaction.
 * @property {string} actionName - The name of the selected action.
 * @property {string} reactionName - The name of the selected reaction.
 *
 */
export const useWebsiteStore = defineStore("websiteStore", {
  state: () => ({
    showNavBar: true,
    showCancelButton: false,
    reactionButtonisDisabled: true,
    showCreateButton: false,
    actionIsSelected: false,
    reactionIsSelected: false,
    actionId: null as string | null,
    actionOptions: {} as Record<string, unknown>,
    reactionId: null as string | null,
    reactionOptions: {} as Record<string, unknown>,
    actionServiceId: null as string | null,
    reactionServiceId: null as string | null,
    actionName: "" as string,
    reactionName: "" as string,
  }),

  actions: {
    loadWorkflowState(this: typeof useWebsiteStore.prototype) {
      const savedState = localStorage.getItem("workflowState");

      if (savedState) {
        try {
          const parsedState = JSON.parse(savedState);
          Object.assign(this, parsedState);

          if (!this.isValidJsonObject(this.actionOptions)) {
            this.actionOptions = {};
          }
          if (!this.isValidJsonObject(this.reactionOptions)) {
            this.reactionOptions = {};
          }
        } catch (err) {
          console.error("Failed to parse saved state", err);
          this.clearWorkflowState();
        }
      }
    },

    saveWorkflowState(this: typeof useWebsiteStore.prototype) {
      const stateToSave = {
        actionId: this.actionId,
        actionOptions: this.actionOptions,
        reactionId: this.reactionId,
        reactionOptions: this.reactionOptions,
        actionServiceId: this.actionServiceId,
        reactionServiceId: this.reactionServiceId,
        showNavBar: this.showNavBar,
        showCancelButton: this.showCancelButton,
        reactionButtonisDisabled: this.reactionButtonisDisabled,
        showCreateButton: this.showCreateButton,
        actionIsSelected: this.actionIsSelected,
        reactionIsSelected: this.reactionIsSelected,
        actionName: this.actionName,
        reactionName: this.reactionName,
      };

      localStorage.setItem("workflowState", JSON.stringify(stateToSave));
    },

    isValidJsonObject(value: unknown): value is Record<string, unknown> {
      return (
        typeof value === "object" && value !== null && !Array.isArray(value)
      );
    },

    clearWorkflowState(this: typeof useWebsiteStore.prototype) {
      this.$reset();
      localStorage.removeItem("workflowState");
      this.saveWorkflowState();
    },

    onActionSelected(this: typeof useWebsiteStore.prototype) {
      this.showNavBar = false;
      this.showCancelButton = true;
      this.reactionButtonisDisabled = false;
      this.actionIsSelected = true;
      this.saveWorkflowState();
    },

    onReactionSelected(this: typeof useWebsiteStore.prototype) {
      this.showCreateButton = true;
      this.reactionIsSelected = true;
      this.reactionButtonisDisabled = false;
      this.saveWorkflowState();
    },

    resetWorkflowPage(this: typeof useWebsiteStore.prototype) {
      this.actionId = null;
      this.reactionId = null;
      this.actionOptions = {};
      this.reactionOptions = {};
      this.actionServiceId = null;
      this.reactionServiceId = null;
      this.showNavBar = true;
      this.showCancelButton = false;
      this.reactionButtonisDisabled = true;
      this.showCreateButton = false;
      this.actionIsSelected = false;
      this.reactionIsSelected = false;
      this.actionName = "";
      this.reactionName = "";
      this.clearWorkflowState();
    },
  },
});
