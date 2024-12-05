<script setup lang="ts">
const username = ref('')
const password = ref('')

const token = useCookie('token')
const loginError = ref<string | null>(null);

interface RegisterResponse {
    token: string;
    message?: string;
}

const apps = ref<string[]>(['i-logos-google-icon', 'i-logos-google-icon', 'i-logos-google-icon']);

const handleLogin = async () => {
    if (!username.value || !password.value) {
        loginError.value = 'Please enter username and password.';
        return;
    }
    try {
        loginError.value = null;

        const response = await $fetch<RegisterResponse>('/api/login', {
            method: 'POST',
            body: {
                username: username.value,
                password: password.value,
            },
        });

        if (response.token) {
            token.value = response.token;
            console.log('Token stored in localStorage:', response.token);
        }
        console.log('Login successful:', response);
        navigateTo('/');
    } catch (error: any) {
        console.error('Login failed:', error);
        loginError.value = error?.data?.message || 'Login failed. Please try again.';
    }
};
</script>

<template>
    <div class="flex items-center justify-center h-screen w-screen bg-custom">
        <UContainer :ui="{ base: 'mx-0', padding: 'p-10', constrained: 'min-w-[30%] max-w-[80%]' }"
            class="bg-custom-section flex flex-col items-center gap-12 border-custom-line border-2 rounded-[3.125rem]">
            <h1 class="text-9xl">Log in</h1>
            <div class="flex flex-col gap-10 min-w-[80%] max-w-[80%]">
                <div class="flex flex-col">
                    <h2 class="text-3xl px-5">Username</h2>
                    <UInput
                        :ui="{ base: 'w-full focus:outline border-2 border-custom-line opacity-100', rounded: 'rounded-[3.125rem]', placeholder: '!px-5 font-light', color: { white: { outline: 'shadow-none bg-custom ring-0' } }, size: { sm: 'text-5xl' } }"
                        v-model="username" />
                </div>
                <div class="flex flex-col">
                    <h2 class="text-3xl px-5">Password</h2>
                    <UInput type="password"
                        :ui="{ base: 'w-full focus:outline border-2 border-custom-line opacity-100', rounded: 'rounded-[3.125rem]', placeholder: '!px-5 font-light', color: { white: { outline: 'shadow-none bg-custom ring-0' } }, size: { sm: 'text-5xl' } }"
                        v-model="password" />
                    <ULink to="/forgotpassword" class="text-xl text-custom-link self-end px-5">Forgot password?</ULink>
                </div>

                <div class="flex flex-col items-center min-w-full">
                    <div v-if="loginError" class="text-red-500 text-xl mb-4">
                        {{ loginError }}
                    </div>
                    <UButton @click="handleLogin" color="black" class="rounded-[3.125rem] min-w-[30%] max-w-[30%]">
                        <p class="text-center text-4xl py-2 min-w-full">Log in</p>
                    </UButton>
                    <p class="text-xl">New? <ULink to="/signup" class="text-custom-link"><u>Sign Up</u></ULink>
                    </p>
                </div>
            </div>
        </UContainer>
    </div>
</template>

<style scoped></style>
