<template>
  <div>
    <h1 v-if="errorMessage">{{ errorMessage }}</h1>
    <p v-else>Downloading your APK...</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      errorMessage: null, // To store the error message
    };
  },
  mounted() {
    const apkUrl = "/apk/client.apk";

    // Check if the file exists
    fetch(apkUrl, { method: "HEAD" })
      .then((response) => {
        if (response.ok) {
          // File exists, trigger download
          const link = document.createElement("a");
          link.href = apkUrl;
          link.download = "client.apk";
          link.click();
        } else {
          // File not found
          this.errorMessage = "File not found";
        }
      })
      .catch(() => {
        // Handle any errors (e.g., network issues)
        this.errorMessage = "File not found";
      });
  },
};
</script>
