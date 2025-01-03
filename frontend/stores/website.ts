export const useWebsiteStore = defineStore("websiteStore", {
  state: () => ({
    showNavBar: true,
    showCancelButton: false,
    reactionButtonisDisabled: true,
    showCreateButton: false,
    actionIsSelected: false,
    reactionIsSelected: false,
    actionId: null as string | null,
    actionOptions: null as unknown | null,
    reactionId: null as string | null,
    reactionOptions: null as unknown | null,
    actionServiceId: null as string | null,
    reactionServiceId: null as string | null,
  }),

  actions: {
    loadWorkflowState() {
      const savedState = localStorage.getItem("workflowState");

      if (savedState) {
        const parsedState = JSON.parse(savedState);
        Object.assign(this, parsedState);
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
      };

      localStorage.setItem("workflowState", JSON.stringify(stateToSave));
      console.log("State saved", stateToSave);
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
      this.actionOptions = null;
      this.reactionOptions = null;
      this.actionServiceId = null;
      this.reactionServiceId = null;
      this.showNavBar = true;
      this.showCancelButton = false;
      this.reactionButtonisDisabled = true;
      this.showCreateButton = false;
      this.actionIsSelected = false;
      this.reactionIsSelected = false;
      this.clearWorkflowState();
    },
  },
});
