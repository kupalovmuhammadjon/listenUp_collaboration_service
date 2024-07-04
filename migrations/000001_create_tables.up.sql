CREATE TYPE collaboration_role AS ENUM ('owner', 'collaborator');
CREATE TYPE invitation_status AS ENUM ('pending', 'accepted', 'declined');

CREATE TABLE collaborations (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    podcast_id uuid not null,
    user_id uuid not null unique,
    role collaboration_role DEFAULT 'collaborator',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE invitations (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    podcast_id uuid not null,
    inviter_id uuid not null,
    invitee_id uuid not null unique,
    status invitation_status DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE comments (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    podcast_id uuid,
    episode_id uuid,
    user_id uuid not null,
    content TEXT not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
