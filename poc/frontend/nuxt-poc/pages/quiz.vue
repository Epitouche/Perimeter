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

const catImageUrl = ref('');

const fetchCatImage = async () => {
    try {
        catImageUrl.value = '';
        const response = await fetch('https://api.thecatapi.com/v1/images/search');
        const data = await response.json();
        catImageUrl.value = data[0].url;
    } catch (error) {
        console.error('Error fetching cat image:', error);
    }
};

const submitQuiz = async () => {
    await fetchCatImage();
};

</script>

<template>
    <div class="flex justify-start items-center pl-10 pt-2">
        <h1 class="text-6xl font-serif">Cat Quiz</h1>
    </div>
    <div class="flex flex-col items-center h-screen gap-20">
        <div>
            <div v-for="(question, index) in questions" :key="question.id" class="flex flex-col">
                <h2 class="text-7xl font-serif pt-20 pb-4">{{ question.text }}</h2>
                <label v-for="option in question.options" :key="option"
                    class="flex flex-row items-center gap-5 text-5xl font-serif pl-10 py-6">
                    <input type="radio" :name="'question-' + question.id" :value="option" class="size-5" />
                    {{ option }}
                </label>
            </div>
        </div>
        <div class="pb-20">
            <UPopover overlay :popper="{ locked: true, placement: 'top-start' }">
                <UButton @click="submitQuiz" class="text-5xl font-serif font-semibold px-6 py-5 rounded-xl">
                    Submit Answers
                </UButton>
                <template #panel>
                    <div v-if="catImageUrl" class="">
                        <img :src="catImageUrl" class="" />
                    </div>
                    <div>
                        <UButton to="/" target="_self" class="text-5xl font-serif font-semibold px-6 py-5 rounded-xl">
                            Try again
                        </UButton>
                    </div>
                </template>
            </UPopover>
        </div>
    </div>

</template>
