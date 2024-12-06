export default defineEventHandler(async (event) => {
    try {
        const params = await readBody(event);
        const { serviceId } = event.context.params;
        const response = await $fetch(`http://server:8080/api/v1/action/info/${serviceId}`, {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer ' + params.token,
            },
        });
        return response;
    } catch (error) {
        console.error('Error fetching services:', error);
        console.log('Error fetching services:', error);
    }
});
