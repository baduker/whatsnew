# whatsnew
a dead simple AWS What's New RSS feed grabber

A super simple RSS feed fetcher for 
[AWS What's new](https://aws.amazon.com/new/?nc1=f_cc&whats-new-content-all.sort-by=item.additionalFields.postDateTime&whats-new-content-all.sort-order=desc) written in Go.

### Usage

1. Clone the repo
2. Build the file `go build -o whatsnew main.go`
3. Run `./whatsnew`

```
1. Amazon Redshift materialized views support external tables | Fri, 19 Jun 2020 21:52:57 +0000
https://aws.amazon.com/about-aws/whats-new/2020/06/amazon-redshift-materialized-views-support-external-tables/

Amazon Redshift adds materialized view support for external tables. With this enhancement, you can create materialized
views in Amazon Redshift that reference external data sources such as Amazon S3 via Spectrum, or data in Aurora or RDS
PostgreSQL via federated queries.

2. NexGuard forensic watermarking is now available with AWS Elemental MediaConvert | Fri, 19 Jun 2020 18:02:29 +0000
https://aws.amazon.com/about-aws/whats-new/2020/06/nexguard-forensic-watermarking-now-available-with-aws-elemental-mediaconvert/

AWS Elemental MediaConvert now supports forensic watermarking using NexGuard. This feature enables you to watermark
content for both mezzanine and OTT streaming contexts in order to enable content leak forensic workflows. You have
access to an added layer of security and traceability for valuable pre-release and early release content, and a simple
way for watermarking during video transcoding and OTT content preparation. NexGuard forensic watermarking in
MediaConvert for pre-release content, including the recently announced NexGuard ClipMark for short form content as well
as NexGuard Streaming for on-demand OTT content, enables full watermark automation when processing in AWS.

...
```