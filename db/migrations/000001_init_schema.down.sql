-- Drop foreign key constraints
ALTER TABLE "profile" DROP CONSTRAINT IF EXISTS "profile_user_id_fkey";
ALTER TABLE "profile" DROP CONSTRAINT IF EXISTS "profile_group_id_fkey";
ALTER TABLE "policy" DROP CONSTRAINT IF EXISTS "policy_group_id_fkey";
ALTER TABLE "permission" DROP CONSTRAINT IF EXISTS "permission_policy_id_fkey";
ALTER TABLE "users_group" DROP CONSTRAINT IF EXISTS "users_group_user_id_fkey";
ALTER TABLE "users_group" DROP CONSTRAINT IF EXISTS "users_group_group_id_fkey";
ALTER TABLE "users" DROP CONSTRAINT IF EXISTS "users_tenant_id_fkey";
ALTER TABLE "group_policy" DROP CONSTRAINT IF EXISTS "group_policy_group_id_fkey";
ALTER TABLE "group_policy" DROP CONSTRAINT IF EXISTS "group_policy_policy_id_fkey";

-- Drop unique indexes
DROP INDEX IF EXISTS "users_id_tenant_id_idx";
DROP INDEX IF EXISTS "permission_id_policy_id_idx";
DROP INDEX IF EXISTS "profile_id_user_id_idx";
DROP INDEX IF EXISTS "group_policy_group_id_policy_id_idx";
DROP INDEX IF EXISTS "users_group_group_id_user_id_idx";

-- Drop tables
DROP TABLE IF EXISTS "users_group";
DROP TABLE IF EXISTS "group_policy";
DROP TABLE IF EXISTS "group";
DROP TABLE IF EXISTS "profile";
DROP TABLE IF EXISTS "permission";
DROP TABLE IF EXISTS "policy";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "tenant";

