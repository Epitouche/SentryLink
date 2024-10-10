<template>
  <div>
    <div class="flex flex-col w-28 bg-cprimary text-white h-screen">
      <ul class="flex flex-col items-center h-full gap-10 pt-10">
        <ULink
          active-class="flex flex-col justify-center items-center rounded-md w-24 h-14 hover:bg-cprimary-dark transition-colors duration-300 text-sky-300"
          inactive-class="flex flex-col justify-center items-center rounded-md w-24 h-14 hover:bg-cprimary-dark transition-colors duration-300"
        >
          <u-icon name="material-symbols:menu" class="w-6 h-6" />
        </ULink>
        <li>
          <ULink
            active-class="flex flex-col justify-center items-center rounded-md w-24 h-14 hover:bg-cprimary-dark transition-colors duration-300 text-sky-300"
            inactive-class="flex flex-col justify-center items-center rounded-md w-24 h-14 hover:bg-cprimary-dark transition-colors duration-300 "
            to="/dashboard"
          >
            <u-icon name="material-symbols:dashboard" class="w-6 h-6" />
            <p>Dashboard</p>
          </ULink>
        </li>
        <li>
          <ULink
            active-class="flex flex-col justify-center items-center rounded-md w-24 h-14 hover:bg-cprimary-dark transition-colors duration-300 text-sky-300"
            inactive-class="flex flex-col justify-center items-center rounded-md w-24 h-14 hover:bg-cprimary-dark transition-colors duration-300"
            to="/settings"
          >
            <u-icon name="material-symbols:settings" class="w-6 h-6" />
            Settings
          </ULink>
        </li>
        <li>
          <ULink
            active-class="flex flex-col justify-center items-center rounded-md w-24 h-14 hover:bg-cprimary-dark transition-colors duration-300 text-sky-300"
            inactive-class="flex flex-col justify-center items-center rounded-md w-24 h-14 hover:bg-cprimary-dark transition-colors duration-300"
            to="/profile"
          >
            <u-icon name="material-symbols:person" class="w-6 h-6" />
            Profile
          </ULink>
        </li>
      </ul>
    </div>
    <NuxtPage />
  </div>
</template>
