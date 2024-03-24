# Route53 DNS record
resource "aws_route53_record" "wendover" {
  zone_id   =   var.web_dns_zone_id
  name      =   var.web_full_domain
  type      =   "CNAME"
  ttl       =   500
  records   =   [aws_cloudfront_distribution.wendover_web.domain_name]
}

# S3 Buckets
resource "aws_s3_bucket" "wendover_logs" {
  bucket = "wendover-logs"
}

resource "aws_s3_bucket_ownership_controls" "wendover_logs" {
  bucket = aws_s3_bucket.wendover_logs.bucket
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_acl" "wendover_logs" {
  depends_on = [aws_s3_bucket_ownership_controls.wendover_logs]

  bucket = aws_s3_bucket.wendover_logs.id
  acl    = "private"
}

resource "aws_s3_bucket" "wendover_web" {
  bucket = var.web_full_domain
}

resource "aws_s3_bucket_ownership_controls" "wendover_web" {
  bucket = aws_s3_bucket.wendover_web.bucket
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_acl" "wendover_web" {
  depends_on = [aws_cloudfront_origin_access_control.wendover_web]

  bucket = aws_s3_bucket.wendover_web.id
  acl    = "private"
}

# CloudFront configuration
resource "aws_cloudfront_origin_access_control" "wendover_web" {
  name                                =   "wendover-web"
  origin_access_control_origin_type   =   "s3"
  signing_behavior                    =   "always"
  signing_protocol                    =   "sigv4"
}

resource "aws_cloudfront_distribution" "wendover_web" {
  origin {
    domain_name              = aws_s3_bucket.wendover_web.bucket_regional_domain_name
    origin_access_control_id = aws_cloudfront_origin_access_control.wendover_web.id
    origin_id                = "wendover-web"
  }

  enabled               =   true
  is_ipv6_enabled       =   true
  default_root_object   =   "index.html"

  logging_config {
    include_cookies   =   false
    bucket            =   aws_s3_bucket.wendover_logs.bucket_regional_domain_name
    prefix            =   "wendover"
  }

  aliases = [var.web_full_domain]

  default_cache_behavior {
    allowed_methods   =   ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods    =   ["GET", "HEAD"]
    target_origin_id  =   "wendover-web"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy  =   "allow-all"
    min_ttl                 =       0
    default_ttl             =    3600
    max_ttl                 =   86400
  }

  # Cache behavior with precedence 0
  ordered_cache_behavior {
    path_pattern     = "/content/immutable/*"
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD", "OPTIONS"]
    target_origin_id = "wendover-web"

    forwarded_values {
      query_string = false
      headers      = ["Origin"]

      cookies {
        forward = "none"
      }
    }

    min_ttl                 =          0
    default_ttl             =      86400
    max_ttl                 =   31536000
    compress                =   true
    viewer_protocol_policy  =   "redirect-to-https"
  }

  # Cache behavior with precedence 1
  ordered_cache_behavior {
    path_pattern      =   "/content/*"
    allowed_methods   =   ["GET", "HEAD", "OPTIONS"]
    cached_methods    =   ["GET", "HEAD"]
    target_origin_id  =   "wendover-web"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    min_ttl                =     0
    default_ttl            =  3600
    max_ttl                = 86400
    compress               = true
    viewer_protocol_policy = "redirect-to-https"
  }

  price_class = "PriceClass_200"

  restrictions {
    geo_restriction {
      restriction_type = "whitelist"
      locations        = ["US"]
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }
}
