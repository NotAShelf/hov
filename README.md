# Hov

**H**yprland **Ov**erview is a tool written in Go that generates a grid of images from your active workspaces, inspired by Sway's [Sov](https://github.com/milgra/sov).
Hov is designed to be used with Hyprland and Hyprland alone. No other window managers are supported.

## Concept

In theory, Hov follows a simple concept: receive necessary information from `hyprctl` and dynamically generate a grid of images using the coordinates of
active windows. In practice, it's slightly more complicated.

### Problem No 1

> ✅ Solved

Hyprctl returns *a lot* of data. We need to agressively parse the json to thin out the data we need. This is done by `get_workspaces()` following a simple
Clients struct.

We filter out `client.Class`, `client.Title`, `client.Address`, `client.Workspaces.Name` and `client.Workspaces.Id` to be able to pinpoint a specific window.
In the future we will also need the Monitor ID to be able to place the window in the correct grid element. That implies the existence of a different grid layout, 
which is not yet implemented nor thought of.

### Problem No 2

> ❌ Not solved (needs help)

The lack of Wayland support in non-gtk or non-qt GUI libraries. I simply do not wish to use GTK or QT bindings. I will personally defenestrate anyone who suggests Rust or
those bindings.

### Problem No 3

> ❌ Not solved (needs help)

Let's assume we have managed to draw the window. How will we place the active windows inside the grid? This requires replicating Hyprland's layout logic, which
actually does differ based on your layout. This is a problem that I have not yet solved.

My current idea is to place scaled rectangles inside each grid element that represent the active window. We can then use the coordinates of the active window
to determine where in a specific grid element active window is.

This will require a lot of math and I am not sure - I suck at math.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
