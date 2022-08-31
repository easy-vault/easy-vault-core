resource "hcp_vault_cluster" "tughi_nirankar_vault_cluster" {
  cluster_id = "tuhi-nirankar"
  hvn_id     = "hvn"
  tier       = "dev"
  public_endpoint = true
}