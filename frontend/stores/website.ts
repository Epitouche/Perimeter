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
    loadWorkflowState() {
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
      console.log("State loaded", this.$state);
    },

    saveWorkflowState() {
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
      console.log("State saved", stateToSave);
    },

    isValidJsonObject(value: unknown): value is Record<string, unknown> {
      return (
        typeof value === "object" && value !== null && !Array.isArray(value)
      );
    },

    clearWorkflowState() {
      this.$reset();
      localStorage.removeItem("workflowState");
      this.saveWorkflowState();
    },

    onActionSelected() {
      this.showNavBar = false;
      this.showCancelButton = true;
      this.reactionButtonisDisabled = false;
      this.actionIsSelected = true;
      this.saveWorkflowState();
    },

    onReactionSelected() {
      this.showCreateButton = true;
      this.reactionIsSelected = true;
      this.reactionButtonisDisabled = false;
      this.saveWorkflowState();
    },

    resetWorkflowPage() {
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
