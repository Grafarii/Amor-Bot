# Models for AI Cosplay Workflow

Download only what your GPU can run. All links are Hugging Face or Civitai — accept licenses before downloading.

## Required (pick ONE path)

### Path A — Full FLUX.1 dev (best quality, ~16 GB VRAM)

| File | Folder | Source |
|------|--------|--------|
| `flux1-dev.safetensors` | `models/diffusion_models/` | [black-forest-labs/FLUX.1-dev](https://huggingface.co/black-forest-labs/FLUX.1-dev) |
| `clip_l.safetensors` | `models/text_encoders/` | same repo, `text_encoder/` |
| `t5xxl_fp16.safetensors` | `models/text_encoders/` | same repo, `text_encoder_2/` |
| `ae.safetensors` | `models/vae/` | same repo, `ae/` |

Use workflow: `workflows/flux-cosplay-base.json` or `flux-cosplay-lora.json`.

**Low VRAM tip:** In UNETLoader set `weight_dtype` to `fp8_e4m3fn` if your build supports it.

### Path B — FP8 checkpoint (simpler, ~12 GB VRAM)

| File | Folder | Source |
|------|--------|--------|
| `flux1-dev-fp8.safetensors` | `models/checkpoints/` | [Comfy-Org/flux1-dev-fp8](https://huggingface.co/Comfy-Org/flux1-dev-fp8) |

Use workflow: `workflows/flux-cosplay-fp8-checkpoint.json`.

### Path C — FLUX.1 Krea dev (best for “not AI-looking” skin)

| File | Folder | Source |
|------|--------|--------|
| `flux1-krea-dev.safetensors` | `models/diffusion_models/` | [black-forest-labs/FLUX.1-Krea-dev](https://huggingface.co/black-forest-labs/FLUX.1-Krea-dev) |
| `clip_l.safetensors` + `t5xxl_fp16.safetensors` + `ae.safetensors` | same as Path A | shared with FLUX dev |

Swap `flux1-dev.safetensors` → `flux1-krea-dev.safetensors` in the base workflow. Same graph, better portrait realism.

---

## Optional LoRAs (cosplay realism)

Place in `models/loras/`. Strength 0.5–0.8 in `flux-cosplay-lora.json`.

Search Civitai for these **categories** (names change; pick highest-rated recent uploads):

| Purpose | Search terms | Typical strength |
|---------|--------------|------------------|
| Photo cosplay look | `flux cosplay realistic`, `flux realism` | 0.6–0.8 |
| Skin texture | `flux skin detail`, `realistic skin flux` | 0.4–0.6 |
| Maid / lace detail | `maid outfit flux`, `lace detail lora` | 0.5–0.7 |
| Chinese dress | `cheongsam flux`, `hanfu cosplay` | 0.5–0.7 |

**Do not stack more than 2 LoRAs** until you know they play well together.

---

## Optional pose control

For extreme poses (contortion, low-angle legs):

1. Install [ComfyUI-ControlNet-Aux](https://github.com/Fannovel16/comfyui_controlnet_aux) (OpenPose preprocessor).
2. For SDXL-based pose workflows, use ControlNet OpenPose — FLUX pose control is still evolving; alternatives:
   - Generate a simpler pose first, then **FLUX Kontext** img2img edit (needs separate workflow).
   - Use **OpenPose reference** in prompt: describe pose explicitly.
   - Try [FLUX ControlNet](https://huggingface.co/XLabs-AI/flux-controlnet-collections) if you have 24 GB+ VRAM.

---

## Optional upscale (post-process)

| Model | Folder | Use |
|-------|--------|-----|
| `4x-UltraSharp.pth` | `models/upscale_models/` | Final 2× upscale |
| `RealESRGAN_x4plus.pth` | `models/upscale_models/` | Alternative upscaler |

Chain after SaveImage: Upscale → VAEDecode (optional second pass at denoise 0.25–0.35).

---

## VRAM cheat sheet

| GPU VRAM | Recommended path |
|----------|------------------|
| 8 GB | FLUX schnell fp8 or GGUF Q4 (needs ComfyUI-GGUF) |
| 12 GB | Path B FP8 checkpoint |
| 16 GB | Path A with fp8 weight dtype |
| 24 GB+ | Path A or C full bf16 + LoRAs |
