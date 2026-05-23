#version 330

// Input vertex attributes
layout (location = 0) in vec3 vertexPosition;
layout (location = 1) in vec2 vertexTexCoord;
layout (location = 2) in vec3 vertexNormal;

// Instancing transformation matrix (Uses 4 vec4 location slots!)
layout (location = 12) in mat4 instanceTransform; 

// Input uniform values
uniform mat4 mvp;

// Output vertex attributes (to fragment shader)
out vec2 fragTexCoord;

void main()
{
    fragTexCoord = vertexTexCoord;
    
    // Project the coordinates using the instance matrix data
    gl_Position = mvp * instanceTransform * vec4(vertexPosition, 1.0);
}