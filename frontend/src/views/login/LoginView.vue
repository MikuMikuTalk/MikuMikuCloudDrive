<script setup lang="ts">
import { ref } from 'vue';

interface LoginForm {
    username: string;
    password: string;
}

const form = ref<LoginForm>({
    username: '',
    password: '',
});

const login = () => {
    // 这里可以添加登录逻辑
    console.log('Logging in with:', form.value);
};
</script>

<template>
    <div class="login-container">
        <h2>Login</h2>
        <form @submit.prevent="login">
            <div class="form-group">
                <label for="username">Username</label>
                <input type="text" id="username" v-model="form.username" placeholder="Enter your username" />
            </div>
            <div class="form-group">
                <label for="password">Password</label>
                <input type="password" id="password" v-model="form.password" placeholder="Enter your password" />
            </div>
            <button type="submit">Login</button>
        </form>
    </div>
</template>

<style scoped>
.login-container {
    max-width: 300px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
}

.form-group {
    margin-bottom: 15px;
}

label {
    display: block;
    margin-bottom: 5px;
}

input {
    width: 100%;
    padding: 8px;
    box-sizing: border-box;
}

button {
    width: 100%;
    padding: 10px;
    background-color: #42b983;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

button:hover {
    background-color: #369f6e;
}
</style>
