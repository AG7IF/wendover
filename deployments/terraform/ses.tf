// Domain Identity
resource "aws_ses_domain_identity" "wendover" {
  domain = var.web_domain
}

resource "aws_route53_record" "wendover_ses_validation" {
  zone_id = var.web_dns_zone_id
  name    = "_amazonses.${aws_ses_domain_identity.wendover.domain}"
  type    = "TXT"
  ttl     = "600"
  records = [aws_ses_domain_identity.wendover.verification_token]
}

resource "aws_ses_domain_identity_verification" "wendover" {
  domain      = aws_ses_domain_identity.wendover.domain
  depends_on  = [aws_route53_record.wendover_ses_validation]
}

// DKIM
resource "aws_ses_domain_dkim" "wendover" {
  domain = aws_ses_domain_identity.wendover.domain
}

resource "aws_route53_record" "wendover_dkim" {
  count   = 3
  zone_id = var.web_dns_zone_id
  name    = "${aws_ses_domain_dkim.wendover.dkim_tokens[count.index]}._domainkey"
  type    = "CNAME"
  ttl     = "600"
  records = ["${aws_ses_domain_dkim.wendover.dkim_tokens[count.index]}.dkim.amazonses.com"]
}

// Email Identity
resource "aws_ses_email_identity" "wendover" {
  email       = "admin@${var.web_domain}"

  depends_on  = [aws_ses_domain_identity.wendover]
}