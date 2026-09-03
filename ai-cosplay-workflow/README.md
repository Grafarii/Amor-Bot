# AI Cosplay Workflow (2.5D Photoreal Style)

ComfyUI workflows and model guide for recreating the **photorealistic anime cosplay** look seen on profiles like [@yunyao114514](https://x.com/yunyao114514): soft window backlight, wooden floors, blue hair, cat ears, maid/cheongsam outfits, low-angle “cosplay photo” framing.

This is a **recreation kit**, not a copy of that creator’s private setup. It uses open FLUX models + optional Civitai LoRAs that match the visual style.

## Quick start

1. Install [ComfyUI](https://github.com/comfyanonymous/ComfyUI) (Desktop or portable).
2. Download the models listed in [models/MODELS.md](models/MODELS.md).
3. Drag a workflow from `workflows/` into ComfyUI.
4. Pick a prompt preset from [prompts/presets.json](prompts/presets.json) and paste it into the positive prompt node.
5. Queue. Tweak seed / guidance / LoRA strength until you like the result.

## Workflows

| File | VRAM | Use when |
|------|------|----------|
| `workflows/flux-cosplay-base.json` | ~16 GB+ | Best quality; separate FLUX dev + CLIP + VAE |
| `workflows/flux-cosplay-fp8-checkpoint.json` | ~12 GB | One FP8 checkpoint; fewer nodes |
| `workflows/flux-cosplay-lora.json` | ~16 GB+ | Base + LoRA slots for cosplay realism |

## Recommended settings

| Setting | Value | Notes |
|---------|-------|-------|
| Resolution | 832×1216 or 1024×1280 | Portrait cosplay framing |
| Steps | 20–28 | FLUX dev; 4 steps if using schnell |
| CFG | **1.0** | Required for FLUX |
| FluxGuidance | 3.0–4.5 | Higher = stronger prompt follow |
| Sampler | euler | scheduler: simple |
| Seed | random | Fix seed when you find a good face |

## Style breakdown (what makes it look like those posts)

- **Base model**: FLUX.1 dev or FLUX.1 Krea dev (natural skin, less “plastic AI look”).
- **Prompting**: full sentences describing a *cosplay photo*, not comma-tag soup.
- **Lighting**: `soft window backlight`, `sheer white curtains`, `warm wooden floor`.
- **Camera**: `low angle`, `35mm lens`, `shallow depth of field`, `DSLR cosplay photo`.
- **Optional LoRAs**: thin realism / cosplay costume LoRAs at 0.5–0.8 strength.
- **Hard poses**: add OpenPose ControlNet (see [models/MODELS.md](models/MODELS.md#optional-pose-control)).

## Folder layout after setup

```
ComfyUI/
├── models/
│   ├── diffusion_models/   # flux1-dev.safetensors (or fp8)
│   ├── text_encoders/      # clip_l.safetensors, t5xxl_fp16.safetensors
│   ├── vae/                # ae.safetensors
│   ├── checkpoints/        # fp8 all-in-one (optional)
│   └── loras/                # optional cosplay LoRAs
└── input/                    # pose reference images (optional)
```

## Troubleshooting

- **Empty CLIP dropdown**: files are in the wrong folder — see MODELS.md paths.
- **CUDA OOM**: use `flux-cosplay-fp8-checkpoint.json` or GGUF quant (ComfyUI-GGUF node).
- **Too anime / not photo enough**: raise guidance to 4.0, add realism LoRA, rewrite prompt as “cosplay photograph”.
- **Bad hands/feet**: lower resolution, add “natural hands” to prompt, or inpaint problem areas.

## License

Workflow JSON is MIT. Model weights follow each model’s own license (FLUX dev = non-commercial unless you have a BFL license).
