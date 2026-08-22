<template>
  <div class="flex min-h-screen flex-col">
    <header class="border-b p-4">Navigation</header>

    <main class="flex-1">
      <slot />
    </main>

    <footer class="border-t p-4 text-sm text-gray-500">Footer</footer>
  </div>
</template>
