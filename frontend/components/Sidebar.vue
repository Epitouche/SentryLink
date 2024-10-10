<template>
  <div>
    <div class="flex flex-col w-28 bg-cprimary text-white h-screen">
      <ul class="flex flex-col items-center h-full gap-10 pt-10">
        <button
          class="flex flex-col justify-center items-center rounded-md w-24 h-14 hover:bg-cprimary-dark hover:text-sky-300"
        >
          <u-icon name="material-symbols:menu" class="w-6 h-6" />
        </button>
        <li>
          <button
            class="flex flex-col justify-center items-center rounded-md w-24 h-14 hover:bg-cprimary-dark hover:text-sky-300"
            to="/"
          >
            <u-icon name="material-symbols:dashboard" class="w-6 h-6" />
            <p>Dashboard</p>
          </button>
        </li>
        <li>
          <button
            class="flex flex-col justify-center items-center rounded-md w-24 h-14 hover:bg-cprimary-dark hover:text-sky-300"
            to="/"
          >
            <u-icon name="material-symbols:settings" class="w-6 h-6" />
            Settings
          </button>
        </li>
        <li>
          <button
            class="flex flex-col justify-center items-center rounded-md w-24 h-14 hover:bg-cprimary-dark hover:text-sky-300"
            to="/"
          >
            <u-icon name="material-symbols:person" class="w-6 h-6" />
            Profile
          </button>
        </li>
      </ul>
    </div>
    <NuxtPage />
  </div>
</template>
