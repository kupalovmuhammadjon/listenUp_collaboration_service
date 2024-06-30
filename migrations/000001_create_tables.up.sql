CREATE TYPE collaboration_role AS ENUM ('owner', 'collaborator', 'viewer');
CREATE TYPE invitation_status AS ENUM ('pending', 'accepted', 'declined');

CREATE TABLE collaborations (
    id uuid PRIMARY KEY,
    podcast_id uuid,
    user_id uuid,
    role collaboration_role DEFAULT 'collaborator',
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE invitations (
    id uuid PRIMARY KEY,
    podcast_id uuid,
    inviter_id uuid,
    invitee_id uuid,
    status invitation_status DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
    id uuid PRIMARY KEY,
    podcast_id uuid,
    user_id uuid,
    content TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
