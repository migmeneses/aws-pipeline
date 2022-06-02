output "instance_id" {
  value = aws_instance.test-instance-flugel.id
}

output "bucket_id" {
  value = aws_s3_bucket.test-bucket-flugel.id
}

