# Models — Best ComfyUI Cosplay Workflow

Download only what your GPU can run. Accept each model’s license before downloading.

## Recommended path (best quality)

### FLUX.1 Krea Dev — best open photoreal base

| File | Folder | Source |
|------|--------|--------|
| `flux1-krea-dev.safetensors` | `models/diffusion_models/` | [black-forest-labs/FLUX.1-Krea-dev](https://huggingface.co/black-forest-labs/FLUX.1-Krea-dev) |
| `clip_l.safetensors` | `models/text_encoders/` | [comfyanonymous/flux_text_encoders](https://huggingface.co/comfyanonymous/flux_text_encoders) |
| `t5xxl_fp16.safetensors` | `models/text_encoders/` | same (use `t5xxl_fp8_e4m3fn.safetensors` under 16 GB VRAM) |
| `ae.safetensors` | `models/vae/` | [black-forest-labs/FLUX.1-dev](https://huggingface.co/black-forest-labs/FLUX.1-dev) `ae.safetensors` |

**Low VRAM tip:** In `UNETLoader` set `weight_dtype` to `fp8_e4m3fn` (already set in the ultimate workflow).

**Even lower VRAM:** use `flux1-krea-dev_fp8_scaled.safetensors` from [Comfy-Org/flux1-krea-dev_fp8_scaled](https://huggingface.co/Comfy-Org/flux1-krea-dev_fp8_scaled) with `flux-krea-fast.json`.

---

## Easy path (~12 GB)

| File | Folder | Source |
|------|--------|--------|
| `flux1-dev-fp8.safetensors` | `models/checkpoints/` | [Comfy-Org/flux1-dev-fp8](https://huggingface.co/Comfy-Org/flux1-dev-fp8) |

Use workflow: `workflows/flux-fp8-simple.json`.

---

## Optional LoRAs (place in `models/loras/`)

Rename to match the workflow dropdowns, or change the node names after load.

| Slot in `flux-krea-lora.json` | Search on Civitai | Strength |
|-------------------------------|-------------------|----------|
| `flux_realism.safetensors` | `flux realism`, `flux skin detail` | 0.5–0.7 |
| `flux_cosplay.safetensors` | `flux cosplay`, `maid outfit flux`, `cheongsam flux` | 0.45–0.65 |

Do not stack more than 2 LoRAs until you know they play well together.

---

## Optional upscaler (`flux-krea-ultimate-upscale.json`)

| File | Folder | Source |
|------|--------|--------|
| `4x-UltraSharp.pth` | `models/upscale_models/` | Search Hugging Face / OpenModelDB for `4x-UltraSharp` |

Alternatives: `RealESRGAN_x4plus.pth`, `4xNomos8kDAT`.

---

## Folder layout after setup

```
ComfyUI/
├── models/
│   ├── diffusion_models/   # flux1-krea-dev.safetensors (or fp8_scaled)
│   ├── text_encoders/      # clip_l + t5xxl_fp16 (or fp8)
│   ├── vae/                # ae.safetensors
│   ├── checkpoints/        # flux1-dev-fp8 (optional simple path)
│   ├── loras/              # optional realism / cosplay LoRAs
│   └── upscale_models/     # optional 4x-UltraSharp
└── input/                  # reference images for img2img
```

## VRAM cheat sheet

| GPU VRAM | Workflow |
|----------|----------|
| 8 GB | FP8 fast + `t5xxl_fp8`, or GGUF via ComfyUI-GGUF |
| 12 GB | `flux-krea-fast.json` or `flux-fp8-simple.json` |
| 16 GB | `flux-krea-ultimate.json` (fp8 weight dtype) |
| 24 GB+ | Ultimate + LoRA + upscale, full bf16 if you prefer |

## Why Krea over plain FLUX.1 Dev?

Krea is tuned to avoid plastic “AI skin”, blown highlights, and generic faces. For cosplay-photo / 2.5D photoreal looks it is currently the best open base that drop-in replaces FLUX.1 Dev in ComfyUI.
