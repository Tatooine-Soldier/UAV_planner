<template>
    <div class="wavy-container">
        <div class="wavy">
            <span>L</span><span>o</span><span>a</span><span>d</span><span>i</span><span>n</span><span>g</span><span>.</span><span>.</span><span>.</span>
        </div>
    </div>
</template>

<style>
    .wavy {
        top: 10%;
        animation: waviy 2.2s infinite;
    }

    @keyframes waviy {
        0%,40%,100% {
          transform: translateY(0)
        }
        20% {
          transform: translateY(-20px)
        }
      }

    .wavy-container {
        animation: fadeIn 2s ease-in-out forwards;
    }

    @keyframes fadeIn {
        from {
          opacity: 0;
        }
        to {
          opacity: 1;
        }
      }
</style>