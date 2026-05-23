#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;

// Input uniform values
uniform sampler2D texture0; // The diffuse texture channel mapped by raylib
uniform vec4 colDiffuse;    // Default base tint color

// Output fragment color
out vec4 finalColor;

void main()
{
    // Fetch texel color from the bound image map and apply base material tint
    finalColor = texture(texture0, fragTexCoord) * colDiffuse;
}