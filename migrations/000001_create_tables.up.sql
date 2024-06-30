CREATE TYPE collaboration_role AS ENUM ('owner', 'collaborator', 'viewer');
CREATE TYPE invitation_status AS ENUM ('pending', 'accepted', 'declined');

CREATE TABLE collaborations (
    id uuid PRIMARY KEY,
    composition_id uuid REFERENCES compositions(id),
    user_id uuid REFERENCES users(id),
    role collaboration_role DEFAULT 'collaborator',
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE invitations (
    id uuid PRIMARY KEY,
    composition_id uuid REFERENCES compositions(id),
    inviter_id uuid REFERENCES users(id),
    invitee_id uuid REFERENCES users(id),
    status invitation_status DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
    id uuid PRIMARY KEY,
    composition_id uuid REFERENCES compositions(id),
    user_id uuid REFERENCES users(id),
    content TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
