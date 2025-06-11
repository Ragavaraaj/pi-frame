# Pi Frame

Pi Frame is a DIY digital photo frame project for Raspberry Pi (or similar Linux SBCs) that displays images from a local folder directly to the framebuffer. It is designed to be simple, reliable, and easy to set up as a kiosk-style photo slideshow.

## Features
- Display images from a local folder
- Supports JPEG and PNG formats
- Automatic image resizing to fit the screen
- Configurable slideshow interval
- Runs directly on the Linux framebuffer (no X11 required)
- Systemd service support for auto-start on boot

## Hardware Requirements
- Raspberry Pi (or any Linux SBC with a framebuffer at `/dev/fb0`)
- Display connected to the device
- Storage for images (SD card, USB, etc.)

## Software Requirements
- Go (for building from source)
- Linux OS with framebuffer support

## Download from Releases (recommended for most users):
   - Pre-built binaries are available on the [GitHub Releases page](https://github.com/Ragavaraaj/pi-frame/releases/latest). Download the latest `photo-frame.zip`, extract it, and use the `photo-frame` binary directly.

## Installation & Build Instructions From source

1. **Clone the repository:**
   ```bash
   git clone <repo-url>
   cd pi-frame
   ```
2. **Build the application (optional if using release):**
   ```bash
   make build
   # The binary will be in build/photo-frame
   ```

## Configuration

Pi Frame can be configured using environment variables:
- `SLIDESHOW_INTERVAL`: Time in seconds between slides (default: 60)
- `SLIDESHOW_DIR`: Directory containing images (default: /home/display/images)

You can set these in a file (e.g., `/etc/default/photoframe`) or directly in your systemd service file.

## Running the Application Locally

To run manually:
```bash
SLIDESHOW_INTERVAL=60 SLIDESHOW_DIR=/home/display/images ./build/photo-frame
```
Or use the provided `make run` for development:
```bash
make run
```

## Systemd Service Setup

To run Pi Frame as a service on boot:
1. Copy the built binary to your desired location (e.g., `/home/display/photo-frame`).
2. Use the provided sample service files in `startup-service/`:
   - `photoframe.service`: Hardcoded environment variables
   - `photoframe.v2.service`: Uses an environment file (recommended)
3. Place your environment file at `/etc/default/photoframe` (see `startup-service/photoframe` for an example).
4. Enable and start the service:
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable photoframe.v2.service
   sudo systemctl start photoframe.v2.service
   ```

## Adding Images
Place your `.jpg`, `.jpeg`, or `.png` images in the directory specified by `SLIDESHOW_DIR`.

## Contributing
Pull requests and issues are welcome! Please open an issue to discuss your ideas or report bugs.

## License
This project is licensed under the MIT License.
