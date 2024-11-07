<script setup lang="ts">
import AOS from "aos";
import "aos/dist/aos.css";

onMounted(() => {
  AOS.init();
});
</script>
<template>
  <div
    class="flex flex-col w-20 h-min p-4 bg-white shadow-md border-2 rounded-3xl border-gray-200 dark:bg-cprimary-dark dark:text-white dark:border-cprimary-dark"
    data-aos="fade-right"
  >
    <ul class="flex flex-col items-center h-full">
      <li>
        <u-link
          active-class="flex justify-center items-center rounded-md p-2 hover:bg-cprimary-light dark:hover:bg-cprimary transition-colors duration-300 black:text-sky-300 text-sky-500"
          inactive-class="flex justify-center items-center rounded-md p-2 hover:bg-cprimary-light dark:hover:bg-cprimary transition-colors duration-300"
          to="/"
        >
          <u-icon name="material-symbols:home-rounded" class="w-10 h-10" />
        </u-link>
      </li>
      <li>
        <u-link
          active-class="flex justify-center items-center rounded-md p-2 hover:bg-cprimary-light dark:hover:bg-cprimary transition-colors duration-300 black:text-sky-300 text-sky-500"
          inactive-class="flex justify-center items-center rounded-md p-2 hover:bg-cprimary-light dark:hover:bg-cprimary transition-colors duration-300 "
          to="/dashboard"
        >
          <u-icon name="material-symbols:dashboard-rounded" class="w-10 h-10" />
        </u-link>
      </li>
      <li>
        <u-link
          active-class="flex justify-center items-center rounded-md p-2 hover:bg-cprimary-light dark:hover:bg-cprimary transition-colors duration-300 black:text-sky-300 text-sky-500"
          inactive-class="flex justify-center items-center rounded-md p-2 hover:bg-cprimary-light dark:hover:bg-cprimary transition-colors duration-300"
          to="/settings"
        >
          <u-icon name="material-symbols:settings-rounded" class="w-10 h-10" />
        </u-link>
      </li>
      <li>
        <u-link
          active-class="flex justify-center items-center rounded-md p-2 hover:bg-cprimary-light dark:hover:bg-cprimary transition-colors duration-300 black:text-sky-300 text-sky-500"
          inactive-class="flex justify-center items-center rounded-md p-2 hover:bg-cprimary-light dark:hover:bg-cprimary transition-colors duration-300"
          to="/profile"
        >
          <u-icon name="material-symbols:person-rounded" class="w-10 h-10" />
        </u-link>
      </li>
    </ul>
  </div>
</template>
