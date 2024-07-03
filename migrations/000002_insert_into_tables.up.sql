INSERT INTO collaborations (id, podcast_id, user_id, role, created_at)
VALUES
    ('1f5e03d7-7e3d-4e02-bbfc-4633ef31f200', NULL, NULL, 'owner', '2024-06-30 12:00:00'),
    ('2c2c76f8-1a4b-4f30-bb0f-83ef9e90b4b0', NULL, NULL, 'collaborator', '2024-06-30 13:00:00'),
    ('3eacdd6d-d52f-463f-b5b7-eb0b2f68b675', NULL, NULL, 'collaborator', '2024-06-30 14:00:00'),
    ('4a265f15-4c95-4b5f-8dbb-1e0c5c98d4a2', NULL, NULL, 'collaborator', '2024-06-30 15:00:00'),
    ('5b9b3c60-35f8-41e3-b38a-5a30839d4304', NULL, NULL, 'owner', '2024-06-30 16:00:00'),
    ('6d9e82fe-dfe0-43d2-ae69-91e32a5e2677', NULL, NULL, 'collaborator', '2024-06-30 17:00:00'),
    ('7fcfa9a1-50cc-4b63-ae85-2d47d3be5e9c', NULL, NULL, 'collaborator', '2024-06-30 18:00:00'),
    ('8e9f4b94-3a5a-4d6f-af20-837a2b0bdf4e', NULL, NULL, 'collaborator', '2024-06-30 19:00:00'),
    ('9a3a62ca-8764-4b58-9a22-d4b8c42be4e9', NULL, NULL, 'owner', '2024-06-30 20:00:00'),
    ('10c346c2-97c4-4461-b9c8-5e7e176f8044', NULL, NULL, 'collaborator', '2024-06-30 21:00:00');

INSERT INTO invitations (id, podcast_id, inviter_id, invitee_id, status, created_at)
VALUES
    ('550e8400-e29b-41d4-a716-446655440000', NULL, NULL, NULL, 'pending', '2024-06-30 08:00:00'),
    ('551e8400-e29b-41d4-a716-446655440001', NULL, NULL, NULL, 'accepted', '2024-06-29 15:30:00'),
    ('552e8400-e29b-41d4-a716-446655440002', NULL, NULL, NULL, 'declined', '2024-06-28 11:45:00'),
    ('553e8400-e29b-41d4-a716-446655440003', NULL, NULL, NULL, 'pending', '2024-06-27 09:20:00'),
    ('554e8400-e29b-41d4-a716-446655440004', NULL, NULL, NULL, 'accepted', '2024-06-26 14:00:00'),
    ('555e8400-e29b-41d4-a716-446655440005', NULL, NULL, NULL, 'pending', '2024-06-25 10:30:00'),
    ('556e8400-e29b-41d4-a716-446655440006', NULL, NULL, NULL, 'declined', '2024-06-24 16:45:00'),
    ('557e8400-e29b-41d4-a716-446655440007', NULL, NULL, NULL, 'pending', '2024-06-23 13:15:00'),
    ('558e8400-e29b-41d4-a716-446655440008', NULL, NULL, NULL, 'accepted', '2024-06-22 09:45:00'),
    ('559e8400-e29b-41d4-a716-446655440009', NULL, NULL, NULL, 'pending', '2024-06-21 11:00:00');
