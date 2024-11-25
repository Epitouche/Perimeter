<script setup lang="ts">

definePageMeta({
    layout: 'default'
});

const questions = [
    {
        id: 1,
        text: 'Which word do you prefer?',
        options: ['Schtroumpfy', 'Caracasieue', 'Prouetz', 'Snerferwerfel']
    },
    {
        id: 2,
        text: 'What cant you live without?',
        options: ['A bowling ball', 'A turtle shell', 'Tom mendy', 'A rotten butternut squash']
    },
    {
        id: 3,
        text: 'Would you rather... ?',
        options: ['Have a hairy tongue', 'Have your sweat smell like cheese', 'Look 90 your whole life', 'Have uncutable 20cm toe nails']
    },
    {
        id: 4,
        text: 'Which is your favorite holiday?',
        options: ['Presidents day', 'Europe day', 'Chinese new year', 'Orthodox St. Cyril and St. Methodius Day']
    },
    {
        id: 5,
        text: 'Which is your favorite animal?',
        options: ['Red-lipped bat fish', 'Naked mole rat', 'Pink fairy armadillo', 'Babirusa']
    }
];

const selectedAnswers = ref(Array(questions.length).fill(''));
const popoverState = ref(false);
const catImageUrl = ref('');

interface CatImage {
  url: string;
}

const fetchCatImage = async () => {
    try {
        catImageUrl.value = '';
        const data = await $fetch<CatImage[]>('https://api.thecatapi.com/v1/images/search');
        catImageUrl.value = data[0]?.url || '';
    } catch (error) {
        console.error('Error fetching cat image:', error);
    }
};

const submitQuiz = async () => {
    if (selectedAnswers.value.includes('')) {
        alert('Answer all questions before submitting');
        popoverState.value = true;
        return;
    }
    await fetchCatImage();
    popoverState.value = true;
};

</script>

<template>
    <div class="flex justify-start items-center pl-10 pt-2">
        <h1 class="text-6xl font-serif">Cat Quiz</h1>
    </div>
    <div class="flex flex-col items-center h-screen gap-20">
        <div>
            <div v-for="(question, index) in questions" :key="question.id" class="flex flex-col">
                <h2 class="text-6xl font-serif pt-20 pb-4">{{ question.text }}</h2>
                <label v-for="option in question.options" :key="option"
                    class="flex flex-row items-center gap-5 text-4xl font-serif py-6">
                    <input type="radio" :name="'question-' + question.id" :value="option"
                        v-model="selectedAnswers[index]" class="size-5" />
                    <p>{{ option }}</p>
                </label>
            </div>
        </div>
        <div class="pb-20">
            <UPopover overlay v-model:open="popoverState" class="UPopover">
                <UButton @click="submitQuiz" class="text-5xl font-serif font-semibold px-6 py-5 rounded-xl">
                    Submit Answers
                </UButton>
                <template #panel>
                    <div class="flex flex-col justify-center items-center bg-black gap-20">
                        <div v-if="catImageUrl">
                            <img :src="catImageUrl" class="rounded-xl " />
                        </div>
                        <div>
                            <UButton to="/" target="_self"
                                class="text-5xl font-serif font-semibold px-6 py-5 rounded-xl">
                                Try again
                            </UButton>
                        </div>
                    </div>
                </template>
            </UPopover>
        </div>
    </div>
</template>

<style scoped>
:deep(.UPopover div.fixed.inset-0.transition-opacity.z-50.bg-gray-200\/75.dark\:bg-gray-800\/75) {
  background-color: black !important;;
}
:deep(.UPopover div.z-50.group) {
    height: 100% !important;
    width: 100% !important;
    top: 50% !important;
    left: 50% !important;
    transform: translate(-50%, -50%) !important;
    z-index: 50 !important;
    display: flex !important;
    justify-content: center !important;
    align-items: center !important;
    align-self: center !important;
}
</style>
