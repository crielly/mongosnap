#!/usr/bin/env python
import requests
import os
import boto3
import shutil

def get_region():
    return requests.get(
        "http://169.254.169.254/latest/dynamic/instance-identity/document/"
    ).json()['region']

def get_ssm_client(region):
    return boto3.client('ssm', region_name=region)

def get_parameter(paramname, client):
    return client.get_parameter(
        Name=paramname
    )['Parameter']['Value']

def get_github_release(owner, repo, tag):
    return requests.get(
        "https://api.github.com/repos/{}/{}/releases/tags/{}".format(
            owner, repo, tag
        )
    ).json()

def del_existing_file(filepath):
    try:
        os.remove(filepath)
    except OSError:
        pass

def save_to_disk(url, diskpath, mode):
    payload = requests.get(
        url, stream=True
    )

    with open(diskpath, 'wb') as outfile:
        shutil.copyfileobj(payload.raw, outfile)
    del payload

    os.chmod(diskpath, mode)

if __name__ == '__main__':
    awsregion = get_region()
    ssmclient = get_ssm_client(awsregion)
    binversion = get_parameter("mongosnapversion", ssmclient)
    release = get_github_release("crielly", "mongosnap", binversion)
    filepath = "/usr/local/bin/mongosnap"

    del_existing_file(filepath)

    save_to_disk(
        release['assets'][0]['browser_download_url'],
        filepath,
        "0555"
    )
