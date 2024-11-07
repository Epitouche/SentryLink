<script setup lang="ts">
import HomeCard from "~/components/HomeCard.vue";
import PriceCard from "~/components/PriceCard.vue";

const InfoCard = {
  title: "Title",
  subtitle: "Subtitle",
  content:
    "Ex fugiat aliquip exercitation ea veniam officia aliqua exercitation dolore non commodo. Ad non incididunt mollit do nisi duis anim do ipsum ex aute enim. Exercitation occaecat sunt aliqua voluptate pariatur. Consectetur culpa quis consequat velit ex enim anim aliquip laboris non Lorem elit quis.",
  image: "https://picsum.photos/id/684/600/400",
};
</script>
<template>
  <div class="flex flex-col items-center w-full gap-24 pt-16">
    <div class="flex justify-center w-full">
      <h1 class="text-7xl font-extrabold">Sentrylink</h1>
    </div>
    <div class="flex flex-row items-center w-1/2 gap-2">
      <u-icon name="material-symbols:search" class="w-6 h-6" />
      <u-input
        class="flex-1"
        color="blue"
        size="xl"
        variant="outline"
        placeholder="Search..."
      />
    </div>
    <HomeCard
      :isReversed="false"
      :title="InfoCard.title"
      :subtitle="InfoCard.subtitle"
      :content="InfoCard.content"
      :image="InfoCard.image"
    />
    <HomeCard
      :isReversed="true"
      :title="InfoCard.title"
      :subtitle="InfoCard.subtitle"
      :content="InfoCard.content"
      :image="InfoCard.image"
    />
    <HomeCard
      :isReversed="false"
      :title="InfoCard.title"
      :subtitle="InfoCard.subtitle"
      :content="InfoCard.content"
      :image="InfoCard.image"
    />
    <div class="flex flex-row gap-10">
      <PriceCard
        plan="Free"
        :price="0"
        :features="['Feature 1', 'Feature 2']"
      />
      <PriceCard class="flex flex-row gap-10"
        plan="Pro"
        :price="30"
        :features="['Feature 1', 'Feature 2', 'Feature 3']"
      />
      <PriceCard class="flex flex-row gap-10"
        plan="Enterprise"
        :price="80"
        :features="['Feature 1', 'Feature 2', 'Feature 3', 'Feature 4']"
      />
    </div>
  </div>
</template>
