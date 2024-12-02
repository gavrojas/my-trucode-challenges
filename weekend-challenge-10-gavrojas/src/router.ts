import Home from '@/components/pages/HomePage.vue';
import Chats from '@/components/pages/ChatPage.vue';

export const routes = [
    { path: '/', component: Home },
    { path: '/kitty/:id', component: Chats, props: true, }
]