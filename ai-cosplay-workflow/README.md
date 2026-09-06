# Best ComfyUI Workflow — FLUX.1 Krea Cosplay Kit

Production-ready ComfyUI graphs for **2.5D photoreal cosplay** (soft window light, wooden floors, maid/cheongsam, anime-inspired faces with real skin).

Built around **FLUX.1 Krea Dev** — currently the best open FLUX base for natural skin and non-plastic photoreal — with a **2-pass latent refine** that beats single-shot generation.

## Start here (pick one)

| Priority | File | VRAM | What you get |
|----------|------|------|----------------|
| **Best** | `workflows/flux-krea-ultimate.json` | ~16 GB | Krea + ModelSamplingFlux → base → **1.5× latent refine** |
| Best (canvas) | `workflows/flux-krea-ultimate-ui.json` | ~16 GB | Same graph with groups + notes on the canvas |
| Best + print | `workflows/flux-krea-ultimate-upscale.json` | ~16 GB | Ultimate + `4x-UltraSharp` → ~2× export |
| LoRA | `workflows/flux-krea-lora.json` | ~16 GB | Ultimate-style dual LoRA (realism + cosplay) |
| Fast | `workflows/flux-krea-fast.json` | ~12 GB | Single-pass Krea FP8 |
| Simple | `workflows/flux-fp8-simple.json` | ~12 GB | One checkpoint, fewest files |
| Img2img | `workflows/flux-krea-img2img.json` | ~12 GB | Refine / restyle a reference photo |

## Quick start

1. Install / update [ComfyUI](https://github.com/comfyanonymous/ComfyUI).
2. Download models from [models/MODELS.md](models/MODELS.md).
3. Drag `workflows/flux-krea-ultimate.json` (or the `-ui` variant) into ComfyUI.
4. Paste a preset from [prompts/presets.json](prompts/presets.json) into the positive prompt.
5. Queue. Keep **CFG = 1.0**. Tune **FluxGuidance** (3.0–4.0) and seed.

## Why this is the “best” workflow

| Technique | Why it matters |
|-----------|----------------|
| **FLUX.1 Krea Dev** | Best open photoreal; less plastic skin than plain FLUX Dev |
| **ModelSamplingFlux** | Correct sigma shift for FLUX resolutions |
| **ConditioningZeroOut** | Proper empty negative for distilled FLUX |
| **2-pass refine** | Pass 1 at 832×1216, Pass 2 at 1.5× with denoise ~0.35 adds detail without nuking composition |
| **Natural-language prompts** | FLUX follows sentences better than SD1.5 tag soup |
| **Optional LoRA / upscale** | Add realism LoRAs or print-size upscale without changing the core graph |

## Recommended settings

| Setting | Ultimate | Fast |
|---------|----------|------|
| Resolution | 832×1216 (portrait) | same |
| Pass 1 steps | 28 | 24 |
| Pass 2 steps | 20 @ denoise 0.35 | — |
| CFG | **1.0** (always) | **1.0** |
| FluxGuidance | 3.5 | 3.5 |
| Sampler / scheduler | euler / simple | euler / simple |

Landscape cosplay floor shots: try **1024×768** (see cheongsam preset).

## Style recipe

- Lighting: `soft window backlight`, `sheer white curtains`, `warm wooden floor`
- Camera: `low angle`, `35mm lens`, `shallow depth of field`, `DSLR cosplay photo`
- Skin: `realistic skin pores`, `natural makeup` — avoid “perfect poreless skin”
- Pose: name joints (`propped on elbows`, `one knee raised`) — avoid vague “contortion”

## Folder layout

```
ComfyUI/
├── models/
│   ├── diffusion_models/   # flux1-krea-dev.safetensors
│   ├── text_encoders/      # clip_l + t5xxl
│   ├── vae/                # ae.safetensors
│   ├── checkpoints/        # optional fp8 all-in-one
│   ├── loras/              # optional
│   └── upscale_models/     # optional 4x-UltraSharp
└── input/                  # img2img references
```

## Troubleshooting

| Problem | Fix |
|---------|-----|
| Empty CLIP dropdown | Wrong folder — see MODELS.md |
| CUDA OOM | Use `flux-krea-fast.json` or fp8 T5; lower to 768×1152 |
| Plastic skin | Stay on Krea; guidance ~3.5; add realism LoRA at 0.5 |
| Extra limbs | Simplify pose wording; use presets that anchor elbows/knees |
| Upscale node missing model | Use `flux-krea-ultimate.json` (no upscaler) or download UltraSharp |
| LoRA file not found | Rename downloads to `flux_realism.safetensors` / `flux_cosplay.safetensors` or edit the node |

## License

Workflow JSON is MIT. Model weights follow each publisher’s license (FLUX Krea/Dev = non-commercial unless you have a BFL license).
