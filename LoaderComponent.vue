<template>
    <div class="wavy">
        <span>L</span><span>o</span><span>a</span><span>d</span><span>i</span><span>n</span><span>g</span><span>.</span><span>.</span><span>.</span>
    </div>
</template>

<style>
    .wavy {
        top: 30%;
        animation: waviy 2s infinite;
    }

    @keyframes waviy {
        0%,40%,100% {
          transform: translateY(0)
        }
        20% {
          transform: translateY(-20px)
        }
      }
</style>