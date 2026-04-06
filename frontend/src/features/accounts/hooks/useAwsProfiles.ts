import { useEffect, useState } from "react";
import { GetAWSProfiles } from "../../../../wailsjs/go/app/App";
import { AWSProfileDTO } from "../types/aws";

export function useAwsProfiles() {
  const [profiles, setProfiles] = useState<AWSProfileDTO[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    GetAWSProfiles()
      .then(setProfiles)
      .finally(() => setLoading(false));
  }, []);

  return { profiles, loading };
}
