<template>
  <section class="panel-appear is-visible delay-2 rounded-[28px] border border-[#E5D8C2] bg-[#FFF9F1] p-6 shadow-[0_26px_90px_-68px_rgba(24,22,20,0.52)] md:p-8">
    <p class="font-mono text-[11px] uppercase tracking-[0.3em] text-[#A16207]">Canales directos</p>
    <div class="mt-5 space-y-3">
      <article v-for="card in contactCards" :key="card.title" class="group rounded-[22px] border border-[#E8DDD0] bg-white/90 p-5 transition-all duration-300 hover:-translate-y-0.5 hover:border-[#CFA24F]">
        <div class="flex items-start gap-4">
          <div class="flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl bg-[#181614] text-[#FFF9F1] transition-colors duration-300 group-hover:bg-[#A16207]">
            <svg v-if="card.icon === 'phone'" viewBox="0 0 24 24" class="h-5 w-5" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
              <path d="M22 16.92v3a2 2 0 0 1-2.18 2A19.8 19.8 0 0 1 3.1 5.18 2 2 0 0 1 5.08 3h3a2 2 0 0 1 2 1.72c.12.89.33 1.76.63 2.58a2 2 0 0 1-.45 2.11L9 10.66a16 16 0 0 0 4.34 4.34l1.25-1.26a2 2 0 0 1 2.11-.45c.82.3 1.69.51 2.58.63A2 2 0 0 1 22 16.92Z" />
            </svg>
            <svg v-else-if="card.icon === 'mail'" viewBox="0 0 24 24" class="h-5 w-5" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
              <rect x="3" y="5" width="18" height="14" rx="2" />
              <path d="m3 7 9 6 9-6" />
            </svg>
            <svg v-else viewBox="0 0 24 24" class="h-5 w-5" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
              <path d="M12 21s6-5.33 6-11a6 6 0 1 0-12 0c0 5.67 6 11 6 11Z" />
              <circle cx="12" cy="10" r="2.25" />
            </svg>
          </div>

          <div class="min-w-0 flex-1">
            <p class="font-mono text-[10px] uppercase tracking-[0.22em] text-[#7A7369]">{{ card.title }}</p>
            <p class="mt-1 font-display text-[clamp(1.15rem,2vw,1.6rem)] font-black uppercase leading-[0.95] text-[#181614]">
              {{ card.value }}
            </p>
            <p class="mt-2 text-[14px] leading-relaxed text-[#5B534A]">{{ card.description }}</p>

            <div class="mt-4 flex flex-wrap items-center gap-3">
              <a
                v-if="card.href"
                :href="card.href"
                class="inline-flex min-h-10 items-center justify-center rounded-full bg-[#181614] px-4 py-2 font-mono text-[10px] uppercase tracking-[0.18em] text-[#FFF9F1] transition-colors duration-300 hover:bg-[#A16207]"
              >
                {{ card.action }}
              </a>
              <button
                v-if="card.title === 'WhatsApp'"
                type="button"
                class="inline-flex min-h-10 items-center justify-center rounded-full border border-[#D6C7B1] px-4 py-2 font-mono text-[10px] uppercase tracking-[0.18em] text-[#181614] transition-colors duration-300 hover:border-[#181614] hover:bg-[#181614] hover:text-[#FFF9F1]"
                @click="emit('copy-phone')"
              >
                {{ copied ? 'Copiado' : 'Copiar' }}
              </button>
            </div>
          </div>
        </div>
      </article>
    </div>
  </section>
</template>

<script setup lang="ts">
defineProps<{
  contactCards: Array<{ title: string; value: string; description: string; href: string | null; action: string; icon: 'phone' | 'mail' | 'pin' }>
  copied: boolean
}>()

const emit = defineEmits<{ (e: 'copy-phone'): void }>()
</script>
